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

-- From PonceER/MyTestRepo --> Add file --> /.github/workflows/MyGoSetup.yml
   1. Event to trigger script -- push and pull_request on [main] branches
   2. Virtual machine OS -- Ubuntu-20.04
   3. Checkout go (i.e. calc) codes
   4. Setup go using Go version 1.13.7 as instructed in activity 2
   5. Install addons/dependencies (see above Addon lists)
  
 -- local repo push the following files: calc.go, calc_test.go to https://github.com/PonceER/MyTestRepo.git
 
 

