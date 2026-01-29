import globals from "globals";
import pluginJs from "@eslint/js";
import tseslint from "typescript-eslint";
import pluginReact from "eslint-plugin-react";
import tailwind from "eslint-plugin-better-tailwindcss";

export default [
  {files: ["**/*.{js,mjs,cjs,ts,jsx,tsx}"]},
  {languageOptions: { globals: globals.browser }},
  pluginJs.configs.recommended,
  ...tseslint.configs.recommended,
  pluginReact.configs.flat.recommended,
  tailwind.configs.recommended,
  {
     settings: {
        react: {
           version: "detect"
        }
     },
     rules: {
         "react/react-in-jsx-scope": "off",
         "react/prop-types": "off"
     }
  },
  {
    ignores: ["**/dist/**", "**/build/**", "**/node_modules/**", "**/_template/**"]
  }
];
