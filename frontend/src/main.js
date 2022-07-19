import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";
import store from "./store";

import PrimeVue from "primevue/config";
import Button from "primevue/button";
import InputText from 'primevue/inputtext';
import Ripple from "primevue/ripple"; //icons

import "primevue/resources/themes/saga-blue/theme.css"; //theme
import "primevue/resources/primevue.min.css"; //core css
import "primeicons/primeicons.css";
import "/node_modules/primeflex/primeflex.css";

const app = createApp(App);

app.use(store);
app.use(router);
app.use(PrimeVue, { ripple: true });

app.component("Button", Button);
app.component("InputText", InputText);
app.directive("Ripple", Ripple);

app.mount("#app");








///estlin uninstalled by command 'npm remove @vue/cli-plugin-eslint' - instal it if needed