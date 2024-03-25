import {  rollupPluginHTML as html } from '@web/rollup-plugin-html';
import resolve from '@rollup/plugin-node-resolve';
import terser from '@rollup/plugin-terser';

export default {
    plugins: [
        html({
            input: 'index.html',
        }),
        // Resolve bare module specifiers to relative paths
        resolve(),
        // Minify JS
        terser({
            ecma: 2021,
            module: true,
            warnings: true,
        }),
    ],
    output: {
        dir: 'dist',
    },
    preserveEntrySignatures: 'strict',
};