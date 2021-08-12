from behave import given, when, then
import functions
from time import sleep

sink_creation_page = 'http://localhost/pages/sinks'

@given("")

@given("User is on the sink management page")
def check_sm_page(context):
    functions.go_to_page(sink_creation_page, context)
    sleep(10)
