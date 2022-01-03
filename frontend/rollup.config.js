import svelte from 'rollup-plugin-svelte';
import commonjs from '@rollup/plugin-commonjs';
import resolve from '@rollup/plugin-node-resolve';
import livereload from 'rollup-plugin-livereload';
import { terser } from 'rollup-plugin-terser';
import sveltePreprocess from 'svelte-preprocess';
import typescript from '@rollup/plugin-typescript';
import css from 'rollup-plugin-css-only';
import replace from "@rollup/plugin-replace"

const production = !process.env.ROLLUP_WATCH;

function serve() {
	let server;

	function toExit() {
		if (server) server.kill(0);
	}

	return {
		writeBundle() {
			if (server) return;
			server = require('child_process').spawn('npm', ['run', 'start', '--', '--dev'], {
				stdio: ['ignore', 'inherit', 'inherit'],
				shell: true
			});

			process.on('SIGTERM', toExit);
			process.on('exit', toExit);
		}
	};
}

function devStagingProd(dev, staging, prod) {
	if (process.env.BUILD == "development") {
		return dev;
	}
	if (process.env.BUILD == "staging") {
		return staging;
	}
	return prod;
}

export default {
	input: 'src/main.ts',
	output: {
		sourcemap: true,
		format: 'iife',
		name: 'app',
		file: 'public/build/bundle.js'
	},
	plugins: [
		replace({
			"http://127.0.0.1:8080": devStagingProd("http://127.0.0.1:8080", "https://backend.staging.gyf.d2a.io", "https://backend.prod.gyf.d2a.io"),
			"127.0.0.1:8080": devStagingProd("127.0.0.1:8080", "backend.staging.gyf.d2a.io", "backend.prod.gyf.d2a.io"),
			"http://localhost:8080": devStagingProd("http://localhost:8080", "https://backend.staging.gyf.d2a.io", "https://backend.prod.gyf.d2a.io"),
			"localhost:8080": devStagingProd("localhost:8080", "backend.staging.gyf.d2a.io", "backend.prod.gyf.d2a.io"),

			"ws://localhost:8080": devStagingProd("ws://localhost:8080", "wss://backend.staging.gyf.d2a.io", "wss://backend.prod.gyf.d2a.io"),
			"ws://127.0.0.1:8080": devStagingProd("ws://127.0.0.1:8080", "wss://backend.staging.gyf.d2a.io", "wss://backend.prod.gyf.d2a.io"),
		}),
		svelte({
			preprocess: sveltePreprocess({ sourceMap: !production }),
			compilerOptions: {
				// enable run-time checks when not in production
				dev: !production
			}
		}),
		// we'll extract any component CSS out into
		// a separate file - better for performance
		css({ output: 'bundle.css' }),

		// If you have external dependencies installed from
		// npm, you'll most likely need these plugins. In
		// some cases you'll need additional configuration -
		// consult the documentation for details:
		// https://github.com/rollup/plugins/tree/master/packages/commonjs
		resolve({
			browser: true,
			dedupe: ['svelte']
		}),
		commonjs(),
		typescript({
			sourceMap: !production,
			inlineSources: !production
		}),

		// In dev mode, call `npm run start` once
		// the bundle has been generated
		!production && serve(),

		// Watch the `public` directory and refresh the
		// browser on changes when not in production
		!production && livereload('public'),

		// If we're building for production (npm run build
		// instead of npm run dev), minify
		production && terser()
	],
	watch: {
		clearScreen: false
	}
};
