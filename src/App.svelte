<script>
  import { onMount } from 'svelte';
  import { Router, Route } from 'svelte-routing';

  import Nav     from './Nav.svelte';
  import Sidebar from './Sidebar.svelte';
  import Home    from './Home.svelte';
  import Steps   from './Steps.svelte';
  import About   from './About.svelte';
  import Terms   from './TermsOfService.svelte';
  import Privacy from './PrivacyPolicy.svelte';
  import Footer  from './Footer.svelte';

  import { theme } from './stores.js';

  onMount(() => {
    const saved = localStorage.getItem('theme') || 'dark';
    theme.set(saved);
    document.documentElement.setAttribute('data-theme', saved);

    theme.subscribe(t => {
      document.documentElement.setAttribute('data-theme', t);
      localStorage.setItem('theme', t);
    });

    console.log('---');
    console.log('Receive these instructions over SSH!');
    console.log('  $ ssh runarelay.org\n');
    console.log('SSH fingerprint: SHA256:Cr/E7i8gvmuHsUUV4GNXSoXBVR1zKC3dSb/k3mGUC+w');
    console.log('---');
  });
</script>

<style>
  .app-shell {
    display: flex;
    flex-direction: column;
    min-height: 100vh;
  }

  .app-body {
    display: flex;
    flex: 1;
    padding-top: var(--nav-h);
  }

  .content-offset {
    margin-left: var(--sidebar-w);
    flex: 1;
    min-width: 0;
    display: flex;
    flex-direction: column;
  }

  @media (max-width: 900px) {
    .content-offset { margin-left: 0; }
  }
</style>

<Router>
  <div class="app-shell">
    <Nav />

    <div class="app-body">
      <Sidebar />

      <div class="content-offset">
        <Route path="/">
          <Home />
        </Route>

        <Route path="/step/:stepName" component={Steps} />

        <Route path="/about" component={About} />

        <Route path="/terms" component={Terms} />

        <Route path="/privacy" component={Privacy} />
      </div>
    </div>

    <Footer />
  </div>
</Router>
