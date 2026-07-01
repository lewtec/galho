/**
 * Centralized error reporting function.
 * Use this function instead of calling console.error or Sentry.captureException directly.
 * All unexpected errors should be funneled through this function.
 */
export function reportError(error: unknown, context?: Record<string, unknown>) {
  // In a real application, you would connect this to an error tracking service like Sentry
  // Sentry.captureException(error, { extra: context });

  // For now, we fallback to console.error to ensure errors are visible during development
  console.error('[ErrorReporter]', error, context ? context : '');
}
