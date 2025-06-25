import http from "k6/http";
import { check, sleep } from "k6";

export let options = {
  vus: 1000, // Number of virtual users
  duration: "30s", // Duration of test
};

const baseUrl = "http://localhost:8080";

export default function () {
  const shortenPayload = JSON.stringify({
    url: "https://example.com",
    customAlias: `ganesh-${__VU}-${__ITER}-${Date.now()}`, // Unique alias for each user
  });

  const headers = { "Content-Type": "application/json" };
  const res = http.post(`${baseUrl}/shorten`, shortenPayload, { headers });

  check(res, {
    "Shorten status was 200": (r) => r.status === 200,
  });

  let shortUrl = null;

  // Only parse if it's JSON
  if (res.headers["Content-Type"]?.includes("application/json")) {
    try {
      const body = JSON.parse(res.body);
      shortUrl = body.shortUrl;
    } catch (e) {
      console.error("Failed to parse JSON:", res.body);
    }
  } else {
    console.error("Unexpected Content-Type:", res.headers["Content-Type"]);
  }

  // Proceed to GET only if shortUrl was parsed
  if (shortUrl) {
    let redirectRes = http.get(shortUrl, {
      redirects: 0, // important: disable automatic redirect following
    });

    check(redirectRes, {
      "Redirect status was 302": (r) => r.status === 302,
      "Redirect location correct": (r) =>
        r.headers.Location === "https://example.com",
    });
      
  }

  sleep(1);
}
