Activity: YAML in CICD
Activity #1: Simple Actions Setup
  see https://github.com/PonceER/HelloWorldTest/blob/main/readme.md
  
Activity #2: Virtual Machine using YML and Github Actions
  see https://github.com/PonceER/HelloWorldTest/blob/main/readme.md

Activity #2(B): Build and Test using YML and GitHub Actions
Add onto the YAML file created in activity #1 and include the relevant tools to create the
following capabilities to the automation.
Note: Created new repo i.e. https://github.com/PonceER/MyTestRepo (this repo) to continue with the exercise/activity

Basic Requirement to add:
Addon | Capablility
3.1. create-issue-action@v1.1 -- To raise notifications in Issues
3.2. github.com/franela/goblin -- To enable test using Goblin
3.3. github.com/tebeka/go2xunit -- To convert report from test to XML
3.4. publish-unit-test-result-action@v1.5 --To read XML and display test results
-- simulated build failure, see https://github.com/PonceER/MyTestRepo/issues/2

-- From PonceER/MyTestRepo --> Add file --> /.github/workflows/myGoSetup.yml
   1. Event to trigger script -- push and pull_request on [main] branches
   2. Virtual machine OS -- Ubuntu-20.04
   3. Checkout go (i.e. calc) codes
   4. Setup go using Go version 1.13.7 as instructed in activity 2
   5. Install addons/dependencies (see above Addon lists)
  
 -- local repo push the following files: calc.go, calc_test.go to https://github.com/PonceER/MyTestRepo.git
    see https://github.com/PonceER/MyTestRepo/commit/4b7de3a2a1b6b8011977c5588d36a60ce94f1b67
 
Activity #2(C): Release/Deployment using YML and GitHub Actions
Add onto the YAML file created in activity #2(B) and include the relevant tools to create the release/ deployment capabilities of your choice to the automation.
-- From PonceER/MyTestRepo --> Add file --> /.github/workflows/myRelease.yml
   1. Event to trigger script -- create and tags
   2. Virtual machine OS -- Ubuntu-20.04
   3. Checkout go (i.e. calc) codes
   4. Validate and Release using docker://goreleaser/goreleaser:latest

A simple tag triggered method will be shown.

-- From PC command line run git tag v2.0, git push --tag


Activity: TDD
See https://github.com/PonceER/WebCalculator/issues/1 for the TDD simulation test and develop with collaboration with gavinerh83.


