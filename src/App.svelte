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

    console.log('SSH: ssh runarelay.org');
  });
</script>

<Router>
  <div class="app-shell">
    <Nav />

    <div class="app-body">
      <Sidebar />

      <div class="content-area">
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
