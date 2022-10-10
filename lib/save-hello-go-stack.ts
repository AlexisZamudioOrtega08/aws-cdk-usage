import * as cdk from '@aws-cdk/core';
import * as lambda from '@aws-cdk/aws-lambda';
import * as dynamodb from '@aws-cdk/aws-dynamodb';
import * as apigw from '@aws-cdk/aws-apigateway';
import * as path from 'path';

export class SaveHelloGoStack extends cdk.Stack {
  constructor(scope: cdk.Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);
    
    // Dynamodb table definition.
    // const grettingsTable = new dynamodb.Table(this, "GreetingsTable", {
    //   partitionKey: {
    //     name: "id", 
    //     type: dynamodb.AttributeType.STRING
    //   }
    // });

    // Lambda function.
    const saveHelloFunction = new lambda.Function(this, "saveHelloFunction", {
      runtime: lambda.Runtime.GO_1_X, 
      handler: "main",
      code: lambda.Code.fromAsset(path.resolve(__dirname, "lambda")),
      // environment: {
      //   GREETINGS_TABLE: grettingsTable.tableName,
      // }
    });

    console.log(lambda.Code.fromAsset(path.resolve(__dirname, "lambda")))

    // Permissions to lambda to dynamo table
    // grettingsTable.grantReadWriteData(saveHelloFunction);

    // Create an API Gateway with one method and path
    const helloAPI = new apigw.RestApi(this, "helloApi");

    // Configure API
    helloAPI.root
      .resourceForPath("hello")
      .addMethod("POST", new apigw.LambdaIntegration(saveHelloFunction))

  }
}
