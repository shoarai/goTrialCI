# goTrialCI
CI trial for Go

![](https://travis-ci.org/shoarai/goTrialCI.svg?branch=master)


## Setup CI for Go by Travis CI
1. Sign up [Travis CI](https://travis-ci.org/).
1. Add a config file.
    ##### travis.yml
    ```yml:.travis.yml
    language: go
    
    go:
      - master  # Latest version
    ```

1. Add a link of build status image to README file.<br>
   The image is updated each time new codes are pushed.
    ##### README.md
    
    ```markdown:README.md
    ![](https://travis-ci.org/"username"/"repository name".svg?branch=master)
    ```
