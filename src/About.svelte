<script>
  import ExternalLink from './ExternalLink.svelte';

  export let location;  // implicitly populated by router
</script>

<style>
.flex-container {
  display: flex;
}

.flex-child {
  flex: 1;
}

.flex-child:first-child {
  margin-right: 48px;
}

.h3-padding {
  padding: 56px 0px 8px 0px;
}

.padding-top {
  margin-top: 104px;
}

div > img:first-child {
  margin-right: 56px;
}

div > img {
  border-radius: 4px;
}

.names {
  font-family: Open Sans;
  font-style: normal;
  font-weight: bold;
  font-size: 24px;
  padding-top: 4px;
  display: flex;
  justify-content: center;
  width: 346px;
}

.names:first-child {
  margin-right: 64px;
}

.align-center {
  display: flex;
  justify-content: center;
  width: 346px;
  padding-top: 8px;
}
</style>


<div class="container-fluid flex-container">

  <div class="flex-child">

    <div class="col-lg-8 col-lg-offset-2">
      <h1 class="padding-top">Creating a better future, together.</h1>
    </div>


    <div class="col-lg-offset-2">
      <div class="col-lg-7">
        <h3 class="col-lg-4 h3-padding">Who we are</h3>
      </div>

      <p class="col-lg-7">It would make sense to introduce in this body of text what this is and how to get started. The rest of this text is</p>

      <div class="col-lg-7"><h3 class="col-lg-4 h3-padding">Our Mission</h3></div>
      
      <p class="col-lg-7">It would make sense to introduce in this body of text what this is and how to get started. The rest of this text is.
      </p>
    </div>

  </div>


  <div class="flex-child">

    <div class="row padding-top">
      
      <div>
        <img src="../img/steve.png" width="346" height="504" alt="Image of Steve Phillips">
        <h5 class="names">Steve Phillips</h5>
        <h6 class="align-center">Founder of Effective.AF</h6>
      </div>

      <div>
        <img src="../img/tim.png" width="346" height="504" alt="Image of Tim Sullivan">
        <h5 class="names">Tim Sullivan</h5>
        <h6 class="align-center">Design Technologist</h6>
      </div>
    </div>

    <!-- second flex child endpoint -->
  </div>

  <!-- container endpoint -->
</div>