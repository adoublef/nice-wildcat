/** @type {import('next').NextConfig} */
const nextConfig = {
    output: 'export',
    
    // Optional: Change links `/me` -> `/me/` and emit `/me.html` -> `/me/index.html`
    // trailingSlash: true,

    // Optional: Prevent automatic `/me` -> `/me/`, instead preserve `href`
    // skipTrailingSlashRedirect: true,

    // Optional: Change the output directory `out` -> `dist`
    // distDir: 'dist',

    // When running Next.js via Node.js (e.g. `dev` mode), proxy API requests
    // to the Go server.
    // rewrites: () => [{ source: "/api", destination: "http://localhost:8080/api" }],
};

module.exports = nextConfig