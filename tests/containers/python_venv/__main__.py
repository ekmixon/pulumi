import pulumi

config = pulumi.Config()
print(f'Hello from {config.require("runtime")}')
