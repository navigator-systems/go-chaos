app = "TestingApp"
description = "this is a test" 


job "script" "python3:script.py" {
    region = "us-west-2"
    namespace = "default"
    config {
        tag = "hello" 
        chaos = "terminate"    
        count = 1
    }
}

notifications "slack" {
    to   = ["C04N5HS91MZ", "C04NJH791MU"]
    body = "chaos experiment running, to check dev teams and get latency in qa"
}

#job "script" "python3:script.py" {
#    region = "us-west-2"
#    namespace = "default"
#    project = "this is a test"
#    config {
#        tag = "hello" 
#        chaos = "terminate"    
#        count = 3
#    }
#}
