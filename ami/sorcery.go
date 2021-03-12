package ami

import (
	"context"
	"strconv"
)

// SorceryMemoryCacheExpire expire (remove) all objects from a sorcery memory cache.
func SorceryMemoryCacheExpire(ctx context.Context, client Client, actionID, cache string) (Response, error) {
	return send(ctx, client, "SorceryMemoryCacheExpire", actionID, map[string]string{
		"Cache": cache,
	})
}

// SorceryMemoryCacheExpireObject expire (remove) an object from a sorcery memory cache.
func SorceryMemoryCacheExpireObject(ctx context.Context, client Client, actionID, cache, object string) (Response, error) {
	return send(ctx, client, "SorceryMemoryCacheExpireObject", actionID, map[string]string{
		"Cache":  cache,
		"Object": object,
	})
}

// SorceryMemoryCachePopulate expire all objects from a memory cache and populate it with all objects from the backend.
func SorceryMemoryCachePopulate(ctx context.Context, client Client, actionID, cache string) (Response, error) {
	return send(ctx, client, "SorceryMemoryCachePopulate", actionID, map[string]string{
		"Cache": cache,
	})
}

// SorceryMemoryCacheStale marks all objects in a sorcery memory cache as stale.
func SorceryMemoryCacheStale(ctx context.Context, client Client, actionID, cache string) (Response, error) {
	return send(ctx, client, "SorceryMemoryCacheStale", actionID, map[string]string{
		"Cache": cache,
	})
}

// SorceryMemoryCacheStaleObject mark an object in a sorcery memory cache as stale.
func SorceryMemoryCacheStaleObject(ctx context.Context, client Client, actionID, cache, object string, reload bool) (Response, error) {
	return send(ctx, client, "SorceryMemoryCacheStaleObject", actionID, map[string]string{
		"Cache":  cache,
		"Object": object,
		"Reload": strconv.FormatBool(reload),
	})
}
