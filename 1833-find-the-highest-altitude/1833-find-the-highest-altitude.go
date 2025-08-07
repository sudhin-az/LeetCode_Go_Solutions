func largestAltitude(gain []int) int {
    currentAltitude := 0
    maxAltitude := 0
    for _, g := range gain {
        currentAltitude += g
        if currentAltitude > maxAltitude {
        maxAltitude = currentAltitude
    }
    }
    return maxAltitude
}