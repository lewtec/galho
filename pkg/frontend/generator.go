package frontend

import (
	"fmt"
	"os"
	"path/filepath"
)

const packageJsonTemplate = `{
  "name": "frontend",
  "module": "index.ts",
  "type": "module",
  "devDependencies": {
    "bun-types": "latest",
    "relay-compiler": "^16.0.0",
    "tailwindcss": "^3.3.0",
    "typescript": "^5.0.0"
  },
  "dependencies": {
    "react": "^18.2.0",
    "react-dom": "^18.2.0",
    "react-relay": "^16.0.0",
    "relay-runtime": "^16.0.0",
    "daisyui": "^3.0.0"
  },
  "peerDependencies": {
    "typescript": "^5.0.0"
  }
}
`

const appTsxTemplate = `import React from 'react';
import { createRoot } from 'react-dom/client';

const App = () => {
  return (
    <div className="p-4">
      <h1 className="text-3xl font-bold underline">
        Hello world!
      </h1>
      <button className="btn btn-primary">Button</button>
    </div>
  );
};

const root = createRoot(document.getElementById('root')!);
root.render(<App />);
`

const tsConfigTemplate = `{
  "compilerOptions": {
    "lib": ["ESNext", "DOM"],
    "module": "esnext",
    "target": "esnext",
    "moduleResolution": "bundler",
    "moduleDetection": "force",
    "allowImportingTsExtensions": true,
    "noEmit": true,
    "composite": true,
    "strict": true,
    "downlevelIteration": true,
    "skipLibCheck": true,
    "jsx": "react-jsx",
    "allowSyntheticDefaultImports": true,
    "forceConsistentCasingInFileNames": true,
    "allowJs": true,
    "types": [
      "bun-types" // add Bun global
    ]
  }
}
`

func Generate(path string) error {
	// Ensure directory exists
	if err := os.MkdirAll(path, 0755); err != nil {
		return err
	}

	// Create package.json
	if err := os.WriteFile(filepath.Join(path, "package.json"), []byte(packageJsonTemplate), 0644); err != nil {
		return err
	}

	// Create App.tsx
	if err := os.WriteFile(filepath.Join(path, "App.tsx"), []byte(appTsxTemplate), 0644); err != nil {
		return err
	}

	// Create tsconfig.json
	if err := os.WriteFile(filepath.Join(path, "tsconfig.json"), []byte(tsConfigTemplate), 0644); err != nil {
		return err
	}

	fmt.Printf("Frontend module generated at %s\n", path)
	return nil
}
