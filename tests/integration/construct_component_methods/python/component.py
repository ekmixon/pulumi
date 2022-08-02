# Copyright 2016-2021, Pulumi Corporation.  All rights reserved.

from typing import Optional

import pulumi

class Component(pulumi.ComponentResource):
    def __init__(self,
                 name: str,
                 first: pulumi.Input[str],
                 second: pulumi.Input[str],
                 opts: Optional[pulumi.ResourceOptions] = None):
        props = {
            "first": first,
            "second": second,
        }
        super().__init__("testcomponent:index:Component", name, props, opts, True)

    @pulumi.output_type
    class GetMessageResult:
        def __init__(self, message: str):
            pulumi.set(self, "message", message)

        @property
        @pulumi.getter
        def message(self) -> str:
            return pulumi.get(self, "message")

    def get_message(self, name: pulumi.Input[str]) -> pulumi.Output['Component.GetMessageResult']:
        __args__ = {'__self__': self, 'name': name}
        return pulumi.runtime.call(
            'testcomponent:index:Component/getMessage',
            __args__,
            res=self,
            typ=Component.GetMessageResult,
        )
