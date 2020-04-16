<script>
  import { Router, Route } from 'svelte-routing';

  import Nav from './Nav.svelte';
  import Home from './Home.svelte';
  import Steps from './Steps.svelte';
  import About from './About.svelte';
  import Terms from './TermsOfService.svelte';
  import Privacy from './PrivacyPolicy.svelte';
  import Footer from './Footer.svelte';
</script>

<Router>

  <Nav />

  <div>
    <Route path="/">
      <Home />
    </Route>

    <Route path="/step/:stepName" component={Steps} />

    <Route path="/about" component={About} />

    <Route path="/terms" component={Terms} />

    <Route path="/privacy" component={Privacy} />

    <!-- Remember to add new URL prefixes to server.go:31 -->
  </div>
  
  <Footer />

</Router>
