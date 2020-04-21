<script>
  import { onMount } from 'svelte';
  import { Router, Route } from 'svelte-routing';

  import Nav from './Nav.svelte';
  import Home from './Home.svelte';
  import Steps from './Steps.svelte';
  import About from './About.svelte';
  import Terms from './TermsOfService.svelte';
  import Privacy from './PrivacyPolicy.svelte';
  import Footer from './Footer.svelte';

  onMount(async () => {
    // Spread the word!
    console.log('---');
    console.log('Receive these instructions over SSH!');
    console.log('  $ ssh runarelay.org\n');
    console.log('SSH fingerprint: SHA256:Cr/E7i8gvmuHsUUV4GNXSoXBVR1zKC3dSb/k3mGUC+w');
    console.log('\nMake sure you have an SSH keypair ($ ls -l ~/.ssh/); create one with\n  $ ssh-keygen -b 4096');
    console.log('---');
  })

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
