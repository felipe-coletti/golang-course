package dog

import "math"

// AgeInDogYears converte anos humanos em anos caninos ajustando por porte.
// dogYears: idade do cão em anos humanos.
// weightKg: peso adulto estimado do cão em kg.
func AgeInDogYears(dogYears float64, weightKg float64) float64 {
    // Etapa 1: idade epigenética
    humanAge := 16*math.Log(dogYears) + 31

    // Etapa 2: ajuste de porte pós 2 anos
    if dogYears > 2 {
        switch {
        case weightKg < 10:
            humanAge *= 0.95
        case weightKg > 25:
            humanAge *= 1.05
        }
    }
    return humanAge
}
