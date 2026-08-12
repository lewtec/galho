import i18n from 'i18next';
import { initReactI18next } from 'react-i18next';
import en from './locales/en.json';
import pt from './locales/pt.json';
import { reportError } from '../utils/reportError';

const resources = {
  en: { translation: en },
  pt: { translation: pt },
};

// Detect browser language
const getBrowserLanguage = (): string => {
  const browserLang = navigator.language.split('-')[0]; // Get language code without region
  // Return browser language if we support it, otherwise fallback to 'en'
  return resources[browserLang as keyof typeof resources] ? browserLang : 'en';
};

// init returns a Promise; surface failures through the shared reporter instead of
// an unhandled rejection (or silent failure if the host swallows them).
void i18n
  .use(initReactI18next)
  .init({
    resources,
    lng: getBrowserLanguage(),
    fallbackLng: 'en',
    interpolation: {
      escapeValue: false,
    },
  })
  .catch((error: unknown) => {
    reportError(error, { source: 'i18n.init' });
  });

export default i18n;
