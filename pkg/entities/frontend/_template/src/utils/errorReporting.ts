export function reportError(error: unknown, context?: Record<string, unknown>) {
  // Centralized error reporting function.
  // In a real application, this should be hooked up to an error tracking service like Sentry.
  console.error("Error reported:", error, context);
}
