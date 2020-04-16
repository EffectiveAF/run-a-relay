const crypto = window.crypto || window.msCrypto;

// Excluding max
//
// Mostly from https://github.com/chancejs/chancejs/issues/232#issuecomment-182500222
function randomIntInRange(min, max) {
  let rand = new Uint32Array(1);
  crypto.getRandomValues(rand);
  const zeroToOne = rand[0] / (0xffffffff + 1);
  return Math.floor(zeroToOne * (max - min)) + min;
}

export function chooseRandomHosts(num) {
  const length = serverHosts.length;
  const goodNdxs = [randomIntInRange(0, length)];

  let maybeGoodNdx;
  while (goodNdxs.length < num) {
    maybeGoodNdx = randomIntInRange(0, length);
    // No dupes allowed!
    if (goodNdxs.indexOf(maybeGoodNdx) === -1) {
      goodNdxs.push(maybeGoodNdx);
    }
  }

  goodNdxs.sort();

  const hosts = goodNdxs.map((ndx) => serverHosts[ndx]);

  return hosts;
}

// Based on
// https://trac.torproject.org/projects/tor/wiki/doc/GoodBadISPs and
// https://community.torproject.org/relay/community-resources/good-bad-isps/
//
// Excludes hosters that are either poorly-rated, have entries not
// updated since January 2017, are reportedly Tor relay-unfriendly,
// has a non-HTTPS website (bad sign!), support colo only, and/or are
// on this list of hosters that The Tor Project says we should avoid
// using:
//
//   OVH SAS (AS16276)
//   Online S.a.s. (AS12876)
//   Hetzner Online GmbH (AS24940)
//   DigitalOcean, LLC (AS14061)
export const serverHosts = [
  {name: 'Creanova', url: 'http://creanova.org/', country: 'Finland'},

  {name: 'myLoc', url: 'https://myloc.de/', country: 'Germany'},
  {name: 'Contabo', url: 'https://contabo.de/', country: 'Germany'},

  {name: 'Trabia', url: 'https://www.trabia.com/', country: 'Moldova'},

  {name: 'LiteServer', url: 'https://www.liteserver.nl/', country: 'Netherlands'},
  {name: 'i3d', url: 'https://www.i3d.net/', country: 'Netherlands'},
  {name: 'KoDDoS', url: 'https://koddos.net/', country: 'Netherlands'},
  {name: 'Worldstream', url: 'https://www.worldstream.nl/', country: 'Netherlands'},

  {name: 'HitMe.pl', url: 'https://hitme.pl/', country: 'Poland'},
  {name: 'Slask DataCenter', url: 'https://sldc.eu/', country: 'Poland'},

  {name: 'MilesWeb', url: 'https://www.milesweb.com/', country: 'Romania'},

  {name: 'Datawagon', url: 'https://datawagon.net/', country: 'US'},
  {name: 'DreamHost', url: 'https://www.dreamhost.com/', country: 'US'},
  {name: 'PulseServers', url: 'https://www.pulseservers.com/', country: 'US'}
];
