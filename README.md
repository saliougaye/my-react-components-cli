# My React Components
My Re-usable components

## How to add a new component?

1. Create a Issue with the name of the components and a description of the component 

2. Start a `feature branch` with the name of the components

3. Create a new folder withe the name of the component in the `components` folder

4. Create a `config.json` (or use the CLI to init a component) like this:

```jsonc
{
    // component name
    "name": "",
    "version": "",
    // dependencies to install for this component
    "dependencies": {}, 
    // dev dependencies to install for this component
    "devDependencies": {}
}
```

5. Create the `src` folder and develop the component in that folder. **IMPORTANT**: every time you add a new dependencies, you must add that in the config.json folder

6. After develop and tested the component, update the changelog with the new version and update the version in the config.json of the component. (this is a fictitious version tagging. this is for to take track to the various changes)

7. Push all to origin and create a pull request from the component feature branch to `develop` with commented the tag of the issue of the component

8. Review the `PR` and if is okay accept

9. After that merge from `develop` to `master` and tag the version like this `<component name>:v<number version>`


## To Do
- Create Workflow
    - Choose a workflow to create a component, structure of the folder ecc.
- Create a CLI (connect to github api)
    - Primary Features
        - Init a component (in this repo)
            - Create component folder in `components`
            - Create src folder
            - Create a `config.json` file
            - Create `CHANGELOG.md` file
            - Run `npm init`
            - Run `npm i -D react typescript @types/react`
        - List the components in the repositories
            - Read the `config.json` file in every component
            - Print the components like this: `<name>@<version>`
        - Download a component (in all react projects)
            - Clone the folder of the component in `src/components/`
            - Install the dependencies listed in config.json
        - .......   
    - 
- .......
