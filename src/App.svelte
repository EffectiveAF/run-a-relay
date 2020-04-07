<script>
  import { Router, Route } from 'svelte-routing';

  import Nav from './Nav.svelte';
  import Home from './Home.svelte';
  import Steps from './Steps.svelte';
  import About from './About.svelte';
</script>

<Router>
  <Nav />
  <div>
    <Route path="/">
      <Home />
    </Route>

    <Route path="/step/:stepName" component={Steps} />

    <Route path="/about" component={About} />
  </div>
</Router>
