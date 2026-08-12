import { useTranslation } from 'react-i18next';

export default function Contact() {
  const { t } = useTranslation();

  return (
    <div className="card bg-base-100 shadow-xl">
      <div className="card-body">
        <h1 className="card-title text-3xl">{t('nav.contact')}</h1>
        <div className="space-y-4">
          <p>{t('contact.intro')}</p>
          <div className="form-control">
            <label className="label">
              <span className="label-text">{t('contact.email')}</span>
            </label>
            <input
              type="email"
              placeholder={t('contact.emailPlaceholder')}
              className="input input-bordered"
            />
          </div>
          <div className="form-control">
            <label className="label">
              <span className="label-text">{t('contact.message')}</span>
            </label>
            <textarea
              className="textarea textarea-bordered h-24"
              placeholder={t('contact.messagePlaceholder')}
            ></textarea>
          </div>
          <button className="btn btn-primary">{t('contact.send')}</button>
        </div>
      </div>
    </div>
  );
}
