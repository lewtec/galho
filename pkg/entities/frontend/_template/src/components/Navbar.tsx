import { Link } from "react-router-dom";
import { useTranslation } from "react-i18next";
import ThemeToggle from "./ThemeToggle";
import LanguageSelector from "./LanguageSelector";

export default function Navbar() {
  const { t } = useTranslation();

  return (
    <nav className="navbar bg-base-100 shadow-lg">
      <div className="navbar-start">
        <div className="dropdown">
          <div tabIndex={0} role="button" className="btn btn-ghost lg:hidden">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              className="h-5 w-5"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
            >
              <path
                strokeLinecap="round"
                strokeLinejoin="round"
                strokeWidth="2"
                d="M4 6h16M4 12h8m-8 6h16"
              />
            </svg>
          </div>
          <ul
            tabIndex={0}
            className="menu menu-sm dropdown-content mt-3 z-[1] p-2 shadow bg-base-100 rounded-box w-52"
          >
            <li>
              <Link to="/">{t("nav.home")}</Link>
            </li>
            <li>
              <Link to="/about">{t("nav.about")}</Link>
            </li>
            <li>
              <Link to="/contact">{t("nav.contact")}</Link>
            </li>
          </ul>
        </div>
        <Link to="/" className="btn btn-ghost text-xl">
          MyApp
        </Link>
      </div>

      <div className="navbar-center hidden lg:flex">
        <ul className="menu menu-horizontal px-1">
          <li>
            <Link to="/">{t("nav.home")}</Link>
          </li>
          <li>
            <Link to="/about">{t("nav.about")}</Link>
          </li>
          <li>
            <Link to="/contact">{t("nav.contact")}</Link>
          </li>
        </ul>
      </div>

      <div className="navbar-end gap-2">
        <LanguageSelector />
        <ThemeToggle />
      </div>
    </nav>
  );
}
