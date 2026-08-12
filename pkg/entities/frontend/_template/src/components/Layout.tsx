import { useEffect } from 'react';
import { Outlet } from 'react-router-dom';
import { useTranslation } from 'react-i18next';
import Navbar from './Navbar';

export default function Layout() {
  const { i18n } = useTranslation();

  // Keep <html lang> aligned with the active locale so assistive tech and
  // browsers do not keep the index.html default after a language switch.
  useEffect(() => {
    const lang = i18n.resolvedLanguage || i18n.language;
    if (lang) {
      document.documentElement.lang = lang;
    }
  }, [i18n.language, i18n.resolvedLanguage]);

  return (
    <div className="min-h-screen bg-base-200">
      <Navbar />
      <main className="container mx-auto px-4 py-8">
        <Outlet />
      </main>
    </div>
  );
}
