from selenium import webdriver
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC
from selenium.webdriver.common.by import By
from selenium.webdriver.support.ui import Select
from selenium.webdriver.common.action_chains import ActionChains
import urllib.request

def go_to_page(page, context, open_browser = 'yes'):
    """Open the page in Chrome browser

    Args:
        page (string): site's URL
        context (behave's object): behave's object that hold contextual information during the running of tests
    """
    page_request = urllib.request.urlopen(page)
    assert page_request.status == 200
    options = webdriver.ChromeOptions()
    if open_browser == "No":
        options.add_argument("--headless")
    options.add_argument('--ignore-ssl-errors=yes')
    options.add_argument('--ignore-certificate-errors')
    options.add_argument('--no-sandbox')
    options.add_argument("--disable-dev-shm-usage")
    context.driver = webdriver.Chrome(options=options)
    context.driver.get(str(page))


def enter_information(html_information_name, information, context, source_by=0):
    """Send information required on a page

    Args:
        html_information_name (string): name of html object
        information (string): information that should be send (for ex: container's id)
        context (behave's object): behave's object that hold contextual information during the running of tests
        source_by (int, optional): type of source by. If 0: source by id; if 1: source by css selector. Defaults to 0.
    """
    if source_by == 0:
        WebDriverWait(context.driver, 15).until(EC.presence_of_element_located((By.ID, html_information_name)))
        data = context.driver.find_element_by_id(str(html_information_name))
        if data.get_attribute("value") != None:
            data.click()
            data.clear()
        data.send_keys(str(information))
    elif source_by == 1:
        WebDriverWait(context.driver, 15).until(EC.presence_of_element_located((By.CSS_SELECTOR, html_information_name)))
        data = context.driver.find_element_by_css_selector(str(html_information_name))
        if data.get_attribute("value") != None:
            data.click()
            data.clear()
        data.send_keys(str(information))
    else:
        raise("source_by not found")

def get_text_from_alert(html_information_name, context,):
    WebDriverWait(context.driver, 15).until(EC.presence_of_element_located((By.CLASS_NAME, html_information_name)))
    data = context.driver.find_element_by_class_name(str(html_information_name))
    return data.text

def press_button(button, context, source_by=0):
    """Press a button on current page

    Args:
        button (string): button's css selector or xpath
        context (behave's object): behave's object that hold contextual information during the running of tests
        source_by (int, optional): type of source by. If 0: source by css selector; if 1: source by xpath. Defaults to 0.
    """
    if source_by==0:
        WebDriverWait(context.driver, 10).until(EC.element_to_be_clickable((By.CSS_SELECTOR, button)))
        context.driver.find_element_by_css_selector(str(button)).click()
    elif source_by==1:
        WebDriverWait(context.driver, 10).until(EC.element_to_be_clickable((By.XPATH, button)))
        context.driver.find_element_by_xpath(str(button)).click()
    else:
        raise("source_by not found")    

def select_drop_down_menu_value(css_selector, text, context):
    values= Select(context.driver.find_element_by_css_selector(str(css_selector)))
    values.select_by_visible_text(str(text))

def click_by_link_text(link_text, context):
    """Click on a link by link text

    Args:
        link_text (string): Text displayed on link
        context (behave's object): behave's object that hold contextual information during the running of tests
    """
    WebDriverWait(context.driver, 10).until(EC.element_to_be_clickable((By.LINK_TEXT, link_text)))
    context.driver.find_element_by_link_text(str(link_text)).click()

def check(page, context):
    """Check if the browser are on the required page and save a screenshot in order to document this

    Args:
        page (string): site's URL
        context (behave's object): behave's object that hold contextual information during the running of tests
    """
    time_stop=0
    while time_stop < 3:        
        try:
            assert (context.driver.current_url == str(page))
            break
        except:
            time_stop = time_stop + 0.1
    assert (context.driver.current_url == str(page))

def get_button_attribute(button_css_selector, attribute, context):
    """Check if a button is enabled on current page

    Args:
        button_css_selector (string): button's css selector in html
        attribute (string): button's attribute wanted
        context (behave's object): behave's object that hold contextual information during the running of tests
    """
    element = context.driver.find_element_by_css_selector(str(button_css_selector))
    return element.get_attribute(str(attribute))

def move_mouse(css_selector, context):
    action = ActionChains(context.driver) # create ActionChains object
    contract =  context.driver.find_element_by_css_selector(str(css_selector))
    action.move_to_element(contract).perform() # move the mouse to the element
    contract.click()
