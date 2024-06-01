# Localization
_Localization module uses `golang.org/x/text/language` and requires language codes to be in BCP 47 tags._

**Aftermath localization is organized in the following way:**
```
- {language_code}/
  - {module_name}.json
```

**Each `json` file is expected to have the following schema:**
```
[
  {
    "key": "required, key_passed_to_printer",
    "value": "required, value returned by the printer",
    "context": "optional, this should add meaningful context to help translators localize the string better"
    "notes": "optional all other notes",
  },
  ...
]
```

## Submitting Issues
When submitting and issue/pull request for a localization change, please include the following for **each change** made:
- path to the file, eg `static/localization/en/discord.json`  
  (_for each key changed_)
  - key
  - why this change was made