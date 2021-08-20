Feature: Sink management



    Scenario Outline: Create sink with "<name_validation>" sink name, "<desc_validation>" description, "<tags_validation>" tags, "<type_validation>" type, "<rh_validation>" remote host and "<usrnm_validation>" username
        Given User is logged in ORB site (email: "<email>", password: "<password>")
             And is on the sink management page
        When Click on "+ NEW SINK" button
            And Enter name label ("<name_label>")
            And Enter descripton ("<description>")
            And Choose sink type
            And Click next button
            And Enter remote host ("<remote_host>")
            And Enter username ("<username>")
            And Enter password ("<sink_password>")
            And Click next button
            And Click save button


        Examples:
            | name_validation | desc_validation | tags_validation | type_validation | rh_validation | usrnm_validation|      email     | password | name_label |    description   | remote_host |   username  |  sink_password |
            |      valid      |       valid     |     valid       |       valid     |     valid     |     valid       | test@email.com | 12345678 |  test_sink | test_create_sink |  localhost  | user_tester | 963852741 |
        #     |      valid      |       valid     |     valid       |       valid     |     valid     |     valid       |
        #     |       any       |        any      |      any        |        any      |       any     |      any        |
        #     |       any       |       valid     |     valid       |       valid     |     valid     |     valid       |
        #     |    invalid      |       valid     |     valid       |       valid     |     valid     |     valid       |
        #     |      valid      |       valid     |     valid       |        any      |     valid     |     valid       |
        #     |      valid      |       valid     |     valid       |    invalid      |     valid     |     valid       |
        #     |      valid      |       valid     |     valid       |       valid     |      any      |     valid       |
        #     |      valid      |       valid     |     valid       |       valid     |    invalid    |     valid       |


    Scenario Outline: Delete sink
        Given User is logged in ORB site (email: "<email>", password: "<password>")
             And is on the sink management page
             And a sink already exist ("<name_label>")
        When Choose "delete sink" option
            And Enter the name of the sink to be deleted ("<sink_name_label>")
            And Press delete button
        Then Referred sink ("<name_label>") should be deleted

        Examples:
            | name_label |      email     | password | sink_name_label |
            |  test_sink | test@email.com | 12345678 |     Value 1     |
