import { createRoot } from 'react-dom/client';
import App from './src/App';
import './index.css';

const el = document.getElementById('root');
if (!el) {
  throw new Error('root element #root not found');
}
createRoot(el).render(<App />);
