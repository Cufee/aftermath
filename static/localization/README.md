# Localization
_Localization module uses `golang.org/x/text/language` and requires language codes to be in BCP 47 tags._

**Aftermath localization is organized in the following way:**
```
- {language_code}/
  - {module_name}.yaml
```

**Each `yaml` file is expected to have the following schema:**
```
- key: internal key for lookups
  value: user facing string
  context (Optional): describe how/where the string will be used
  notes (Optional) any additional notes
...
```

## Submitting Issues
When submitting and issue/pull request for a localization change, please include the following for **each change** made:
- path to the file, eg `static/localization/en/discord.yaml`  
  (_for each key changed_)
  - key
  - why this change was made