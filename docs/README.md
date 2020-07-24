# Example::Time::Resource

This resource test async call.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "Type" : "Example::Time::Resource",
    "Properties" : {
        "<a href="#sleep" title="Sleep">Sleep</a>" : <i>Double</i>,
        "<a href="#delay" title="Delay">Delay</a>" : <i>Double</i>,
        "<a href="#enviroment" title="Enviroment">Enviroment</a>" : <i>String</i>
    }
}
</pre>

### YAML

<pre>
Type: Example::Time::Resource
Properties:
    <a href="#sleep" title="Sleep">Sleep</a>: <i>Double</i>
    <a href="#delay" title="Delay">Delay</a>: <i>Double</i>
    <a href="#enviroment" title="Enviroment">Enviroment</a>: <i>String</i>
</pre>

## Properties

#### Sleep

function sleep time in seconds

_Required_: Yes

_Type_: Double

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Delay

Callback delay in seconds

_Required_: No

_Type_: Double

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Enviroment

The Enviroment of the resource

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

## Return Values

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html).

#### PDID

Unique identifier

