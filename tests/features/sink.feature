Feature: Sink management


    Scenario Outline: Create sink with "<name_validation>" sink name, "<desc_validation>" description, "<tags_validation>" tags, "<type_validation>" type, "<rh_validation>" remote host and "<usrnm_validation>" username
        Examples:
            | name_validation | desc_validation | tags_validation | type_validation | rh_validation | usrnm_validation|
            |      valid      |       valid     |     valid       |       valid     |     valid     |     valid       |
            # |      valid      |        any      |      any        |       valid     |     valid     |      any        |
            # |       any       |        any      |      any        |        any      |       any     |      any        |
            # |       any       |       valid     |     valid       |       valid     |     valid     |     valid       |
            # |    invalid      |       valid     |     valid       |       valid     |     valid     |     valid       |
            # |      valid      |       valid     |     valid       |        any      |     valid     |     valid       |
            # |      valid      |       valid     |     valid       |    invalid      |     valid     |     valid       |
            # |      valid      |       valid     |     valid       |       valid     |      any      |     valid       |
            # |      valid      |       valid     |     valid       |       valid     |    invalid    |     valid       |
        # Given User is logged in ORB site
        #     And is on the sink management page
        # When Enter "<type1>" email address ("<email>")
        #     And Enter "<type2>" password ("<password>")
        #     And Press the ORB button Log In
        # Then Access to ORB "<mode>" be enabled


        # Examples:
        #     | name_validation | desc_validation | tags_validation | type_validation | rh_validation | usrnm_validation|
        #     |      valid      |        any      |      any        |       valid     |     valid     |      any        |
        #     |      valid      |       valid     |     valid       |       valid     |     valid     |     valid       |
        #     |       any       |        any      |      any        |        any      |       any     |      any        |
        #     |       any       |       valid     |     valid       |       valid     |     valid     |     valid       |
        #     |    invalid      |       valid     |     valid       |       valid     |     valid     |     valid       |
        #     |      valid      |       valid     |     valid       |        any      |     valid     |     valid       |
        #     |      valid      |       valid     |     valid       |    invalid      |     valid     |     valid       |
        #     |      valid      |       valid     |     valid       |       valid     |      any      |     valid       |
        #     |      valid      |       valid     |     valid       |       valid     |    invalid    |     valid       |


