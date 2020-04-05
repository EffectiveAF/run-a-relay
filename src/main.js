import App from './App.svelte';

const app = new App({
  target: document.body,
  props: {
    appName: 'Run A Relay'
  }
});

export default app;
