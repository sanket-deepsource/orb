from behave import given, when, then
import functions
from time import sleep

login_page = "http://localhost:80/auth/login"
all_sinks_page = 'http://localhost/pages/sinks'
sink_add_page = 'http://localhost/pages/sinks/add'

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
    context.popup_messsage = functions.get_toastr_message("/html/body/ngx-app/ngx-pages/ngx-one-column-layout/nb-layout/div[2]/div/div/nb-toastr-container/nb-toast", context)

@then('A new sink ("{result}") be created')
def check_sink(context, result):
    if result == 'should':
        assert ("Sink sucessfully created" in context.popup_messsage)
        functions.check(all_sinks_page, context)
    elif result == "shouldn't":
        assert ("Failed to create Sink" in context.popup_messsage)
        functions.check(sink_add_page, context)
    else:

        raise("Invalid result")  

#------------delete_sink----------------


@given('a sink already exist ("{name_label}")')
def existing_sink(context, name_label):
    assert len(functions.get_sinks_name(context)) > 0


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
    functions.press_button(".del-button", context)

@when('Enter the name of the sink to be deleted ("{name_label}")')
def name_label_to_be_deleted(context, name_label):
    functions.enter_information("input.size-medium", name_label, context, 1)

@when("Press delete button")
def delete_button(context):
    functions.press_button(".appearance-filled", context)

@then('Referred sink ("{name_label}") should be deleted')
def check_sink_deleted(context, name_label):
    pass


#--------cancel_delete_sink-----------

@when("Press 'CLOSE' button")
def cancel_delete_sink(context):
   functions.press_button(".orb-close-dialog", context) 


#--------detail------------------
@when('Press visualize sink button')
def visualize_sink(context):
    functions.move_mouse('.orb-action-hover-container', context)
    #functions.press_button(".detail-button", context)
    sleep(5)


@then('Sink details page should be displayed')
def details_page(context):
    functions.find_elements_by_class("cdk-overlay-backdrop-showing",context)
    sleep(2)
