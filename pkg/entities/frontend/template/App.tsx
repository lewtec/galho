import React from 'react';
import { createRoot } from 'react-dom/client';
import './index.css';

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
