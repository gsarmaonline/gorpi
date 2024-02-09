# Gorpi REST API integration

The purpose of this module is to provide the fastest
way to declare RESTful APIs. Motivation is to provide
something similar to Ruby on Rails scaffolding.

There are 2 approaches that we can take here.
1. Provide a scaffolding based approach similar to Rails
2. Provide a framework which pre-defines the functions common
to the APIs.


Few opinions:
1. Scaffolding results in a higher amount of code being generated
regardless.
2. Since the code is being generated, it also
provides a higher level of control once the code is generated.
3. It is difficult to control/modify the behaviour of the scaffolded
code once generated.
