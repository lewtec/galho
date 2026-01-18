import js from "@eslint/js";
import tseslint from "typescript-eslint";
import react from "eslint-plugin-react";
import reactHooks from "eslint-plugin-react-hooks";
import betterTailwind from "eslint-plugin-better-tailwindcss";
import globals from "globals";

export default tseslint.config(
  {
    ignores: ["**/dist", "**/node_modules", "**/*.tmpl", "**/*.go", "**/.mise"],
  },
  js.configs.recommended,
  ...tseslint.configs.recommended,
  betterTailwind.configs.recommended,
  {
    files: ["**/*.{js,jsx,ts,tsx}"],
    languageOptions: {
      ecmaVersion: 2020,
      globals: {
        ...globals.browser,
        ...globals.node,
      },
      parserOptions: {
        ecmaFeatures: { jsx: true },
      },
    },
    plugins: {
      react,
      "react-hooks": reactHooks,
    },
    rules: {
      ...react.configs.recommended.rules,
      ...reactHooks.configs.recommended.rules,
      "react/react-in-jsx-scope": "off",
      "@typescript-eslint/no-explicit-any": "warn",
      "@typescript-eslint/no-unused-vars": "warn",

      // Disable troublesome tailwind rules for templates
      "better-tailwindcss/no-unknown-classes": "off",
      "better-tailwindcss/enforce-consistent-line-wrapping": "off",
    },
    settings: {
      react: { version: "detect" },
    },
  },
  {
    files: ["**/*.js"],
    rules: {
      "@typescript-eslint/no-require-imports": "off",
      "@typescript-eslint/no-var-requires": "off"
    }
  }
);
