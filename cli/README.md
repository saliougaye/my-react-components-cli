# CLI


## Commands
All only interactive


### Key Commands
- init
- component init (**only on this repo**)
- component completed (**only on this repo**)
- component ls
- component download

### Feature Commands

#### init
Example: 
```cmd
> ./cmd init
  Open github for authentication.....
  Authenthicated correctly
```

**Notes**
- Check if is me and only me



#### component init

Example: 
```cmd
> ./cmd component init
  - Name of the component: Component Name

  Creating Issue for `Component Name`..... 
  Issue Created -> #`Issue id`

  Creating feature/`Component Name`  branch...
  feature/`Component Name` created

  Creating `Component Name` folder structure
    `Component Name` folder created
    src folder created
    test_env created
    CHANGELOG.md created
    README.md created
    config.json created
    
  Component Name Initialized Correctly ✔

```
**Notes**
- Check if is the my-react-components repo

#### component completed

Example: 
```cmd
> ./cmd component completed
  Component Name completed? (Y/n): y
  config.json updated with the new version and the dependencies? (Y/n): y
  CHANGELOD.md updated with the new version? (Y/n): y
  Are you sure? (Y/n): y

  git add..... ✔
  git commit..... ✔ 
  Pushing all to origin..... ✔
  Creating PR to develop..... ✔
  Accepting PR..... ✔
  Merge develop in master..... ✔
  Tag with version `<component name>:v<number version>`..... ✔
  
  Component Name Added Successfully

```

**Notes**
- Check if is the my-react-components repo


#### component ls

Example: 
```cmd
> ./cmd component ls
  Total Components: 4
  component1@v.1.0.0
  component2@v.1.0.0
  component3@v.1.0.0
  component4@v.1.0.0

```

#### component download

Example: 
```cmd
> ./cmd component download
  > Select component: component1@v.1.0.0
    component1@v.1.0.0
    component2@v.1.0.0
    component3@v.1.0.0
    component4@v.1.0.0


  Cloning folder in src/components..... ✔
  Installing dependencies..... ✔

  component1@v.1.0.0 added correctly ✔


```
