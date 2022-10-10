# SAVE HELLO GO

This is an example of the use of cdk toolkit with lambda functions

## Steps to deploy a project through cdk
* `create your project` create the project you want to deploy
* `modify the stack` under lib folder modify the ts file to add features to your stack (dynamodb, lambda ...), here you need to add the paths to your project.
* `npm run build` create the project you want to deploy
* `cdk deploy`      deploy this stack to your default AWS account/region

## Useful commands

* `npm run build`   compile typescript to js
* `npm run watch`   watch for changes and compile
* `npm run test`    perform the jest unit tests
* `cdk deploy`      deploy this stack to your default AWS account/region
* `cdk diff`        compare deployed stack with current state
* `cdk synth`       emits the synthesized CloudFormation template
