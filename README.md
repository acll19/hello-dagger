# hello-dagger

This is an example application for use with the Dagger Quickstart. It uses the Vue 3 + Vite template, with minor modifications.

> [!NOTE]  
> The dagger-node directory is a [Git submodule](https://git-scm.com/book/en/v2/Git-Tools-Submodules) to a repository with the same name containing the both dagger directory and the dagger.json file. This is an attempt to provide reusability in daggerlike pipelines. See the dagger-run.sh and dagger-publish.sh for more details. Keep in mind that this is a prove of concept and it may lack flexibility.


## Project Setup

```sh
npm install
```

### Compile and Hot-Reload for Development

```sh
npm run dev

# or with dagger
./dagger-run.sh . dev
```

### Type-Check, Compile and Minify for Production

```sh
npm run build

# or with dagger
./dagger-run.sh . build
```

### Run Unit Tests with [Vitest](https://vitest.dev/)

```sh
npm run test:unit

# or with dagger
./dagger-run.sh . test:unit
```

### Run End-to-End Tests with [Cypress](https://www.cypress.io/)

```sh
npm run test:e2e:dev

# or with dagger
./dagger-run.sh . test:e2e:dev
```

This runs the end-to-end tests against the Vite development server.
It is much faster than the production build.

But it's still recommended to test the production build with `test:e2e` before deploying (e.g. in CI environments):

```sh
npm run build
npm run test:e2e

# or with dagger
./dagger-build-e2e.sh .
```

### Lint with [ESLint](https://eslint.org/)

```sh
npm run lint

# or with dagger
./dagger-run.sh . lint
```
