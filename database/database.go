package main

// Add a new secret
func AddSecret(
    user User,
    name string,
    value string,
    password string,
    allowedUses int8,
    lifetime int64,
) {}

// Remove a secret and references from all tables
func DestroySecret(
    name string,
    password string,
) {}

// Get a secret
func FetchSecret(
    name string,
    user User
) {}

// Get a secret with password
func FetchSecretWithPassword() {}

// Add a specific permission for a user
func AddPermission() {}

// Remove a specific permission from a user
func RemovePermission() {}

// Add new user
func AddUser() {}

// Remove user references from all tables 
func RemoveUser() {}

// Remove all permissions for a given user, and set admin
func RemoveUserPermissions() {}

// Make user admin
func PromoteUser() {}

// Remove user admin
func DemoteUser() {}
