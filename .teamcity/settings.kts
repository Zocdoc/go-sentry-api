import jetbrains.buildServer.configs.kotlin.*
import jetbrains.buildServer.configs.kotlin.buildFeatures.perfmon
import jetbrains.buildServer.configs.kotlin.buildSteps.script
import jetbrains.buildServer.configs.kotlin.triggers.vcs

version = "2023.11"

project {
    description = "Go Project Build Configuration"

    // Define Go versions to test against
    val goVersions = listOf("1.8", "1.9", "1.10", "latest")

    // Create a build type for each Go version
    goVersions.forEach { goVersion ->
        buildType {
            id("Go_${goVersion.replace(".", "_")}")
            name = "Build - Go $goVersion"

            // VCS Root
            vcs {
                root(DslContext.settingsRoot)
            }

            // Environment variables
            params {
                param("env.GOPATH", "%system.teamcity.build.workingDir%/go")
                param("env.PATH", "%env.GOPATH%/bin:%env.PATH%")
                param("env.GO_VERSION", goVersion)
            }

            // Build steps
            steps {
                // 1. Install Go version
                script {
                    name = "Install Go"
                    scriptContent = """
                        #!/bin/bash
                        if [ "%env.GO_VERSION%" == "latest" ]; then
                            GO_VERSION="tip"
                        else
                            GO_VERSION="%env.GO_VERSION%"
                        fi
                        
                        echo "Installing Go ${'$'}GO_VERSION"
                        # Add Go installation commands here
                    """.trimIndent()
                }

                // 2. Get golint
                script {
                    name = "Install golint"
                    scriptContent = "go get -u github.com/golang/lint/golint"
                }

                // 3. Run golint
                script {
                    name = "Run golint"
                    scriptContent = "golint ./..."
                }

                // 4. Run tests
                script {
                    name = "Run tests"
                    scriptContent = "go test -v"
                }

                // 5. Install dependencies
                script {
                    name = "Install dependencies"
                    scriptContent = "go get -v -t ."
                }
            }

            // Allow failures for 'latest' version
            if (goVersion == "latest") {
                allowExternalStatus = true
                failureConditions {
                    // Don't fail build if tests fail on latest/tip version
                    executionTimeoutMin = 30
                }
            }

            // Build Features
            features {
                perfmon {}
            }

            // Build Triggers
            triggers {
                vcs {
                    branchFilter = "+:*"
                }
            }

            requirements {
                // Ensure Go is available on the agent
                contains("env.GO_INSTALLED", "true")
            }
        }
    }

    // Build Chain Configuration
    buildTypesOrderBy {
        // Run stable versions first, latest/tip last
        order {
            goVersions.filterNot { it == "latest" }.forEach {
                buildType("Go_${it.replace(".", "_")}")
            }
            buildType("Go_latest")
        }
    }
}
