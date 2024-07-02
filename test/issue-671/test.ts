import { assertStringIncludes } from "jsr:@std/assert";

Deno.test("issue #671", async () => {
  const res = await fetch(
    "http://localhost:8080/flowbite-react@v0.4.9?alias=react:preact/compat,react-dom:preact/compat",
  );
  await res.body?.cancel();
  const id = res.headers.get("x-esm-path");
  const code = await fetch(
    `http://localhost:8080/${id}`,
  ).then((res) => res.text());
  assertStringIncludes(code, "compat/jsx-runtime.js");
});
