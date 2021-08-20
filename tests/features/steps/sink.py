from behave import given, when, then
import functions
from time import sleep

login_page = "http://localhost:80/auth/login"
sink_creation_page = 'http://localhost/pages/sinks'

@given('User is logged in ORB site (email: "{email}", password: "{password}")')
def check_sm_page(context, email, password):
    functions.go_to_page(login_page, context)
    functions.enter_information('input-email', email, context, source_by=0)
    functions.enter_information('input-password', password, context, source_by=0)
    functions.press_button('.appearance-filled', context)

@given("is on the sink management page")
def sink_management_page(context):
    functions.press_button(".eva-settings-2-outline", context)
    functions.click_by_link_text("Sink Management", context)

@when('Click on "+ NEW SINK" button')
def create_new_sink(context):
    functions.press_button(".button-new-sink", context)

@when('Enter name label ("{name_label}")')
def enter_label_name(context, name_label):
    functions.enter_information("nb-form-field.nb-transition:nth-child(1) > div:nth-child(1) > div:nth-child(1) > input:nth-child(2)", name_label, context, 1)

@when('Enter descripton ("{description}")')
def enter_description(context, description):
    functions.enter_information("nb-form-field.nb-transition:nth-child(2) > div:nth-child(1) > input:nth-child(2)", description, context, 1)

@when('Choose sink type')
def sink_type(context):
    functions.press_button(".select-button", context)
    functions.press_button('/html/body/ngx-app/ngx-pages/ngx-one-column-layout/nb-layout/div[2]/div/div/nb-option-list/ul', context, 1)
    #context.driver.find_element_by_xpath(str('/html/body/ngx-app/ngx-pages/ngx-one-column-layout/nb-layout/div[2]/div/div/nb-option-list/ul')).click()

@when('Click next button')
def click_next(context):
    functions.press_button(".next-button", context)

@when('Enter remote host ("{remote_host}")')
def remote_host(context, remote_host):
    functions.enter_information("nb-form-field.ng-star-inserted:nth-child(1) > div:nth-child(1) > div:nth-child(2) > input:nth-child(1)", remote_host, context, 1)    

@when('Enter username ("{username}")')
def username(context, username):
    functions.enter_information("nb-form-field.ng-star-inserted:nth-child(2) > div:nth-child(1) > div:nth-child(2) > input:nth-child(1)", username, context, 1)    

@when('Enter password ("{sink_password}")')
def password(context, sink_password):
    functions.enter_information("nb-form-field.ng-star-inserted:nth-child(3) > div:nth-child(1) > div:nth-child(2) > input:nth-child(1)", sink_password, context, 1)    


@when('Click save button')
def click_save(context):
    functions.press_button(".next-button", context)    

@then('A new sink should be created')
def check_sink(context):
    pass

#----------------------------


@given('a sink already exist ("{name_label}")')
def existing_sink(context, name_label):
    pass


@when('Choose "delete sink" option')
def choose_delete_sink(context):
    # functions.go_to_page(login_page, context)
    # functions.enter_information('input-email', email, context, source_by=0)
    # functions.enter_information('input-password', password, context, source_by=0)
    # functions.press_button('.appearance-filled', context)
    # functions.press_button(".eva-settings-2-outline", context)
    # functions.click_by_link_text("Sink Management", context)
    pass


    # functions.press_button(".button-new-sink", context)
    functions.move_mouse('.orb-action-hover-container', context)
    
    sleep(4)

@when('Enter the name of the sink to be deleted ("{name_label}")')
def name_label_to_be_deleted(context, name_label):
    pass

@when("Press delete button")
def delete_button(context):
    pass

@then('Referred sink ("{name_label}") should be deleted')
def check_sink_deleted(context, name_label):
    pass
