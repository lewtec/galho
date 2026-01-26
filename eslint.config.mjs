import js from "@eslint/js";
import tseslint from "typescript-eslint";
import tailwind from "eslint-plugin-better-tailwindcss";
import globals from "globals";

export default tseslint.config(
  {
    ignores: [
      "dist/**",
      "bin/**",
      "**/node_modules/**",
      "**/.git/**",
      "**/*.d.ts",
      "go.sum",
      "go.mod",
      "**/*.go",
      "**/_template/**",
      "**/debug_eslint.mjs",
      "install_mise.sh",
    ],
  },
  js.configs.recommended,
  ...tseslint.configs.recommended,
  {
    files: ["**/*.{js,mjs,cjs,ts,jsx,tsx}"],
    languageOptions: {
      globals: {
        ...globals.browser,
        ...globals.node,
      },
    },
    plugins: {
      "better-tailwindcss": tailwind,
    },
    rules: {
      ...tailwind.configs.recommended.rules,
    },
  },
);
