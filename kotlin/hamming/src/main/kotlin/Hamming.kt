object Hamming {

    fun compute(leftStrand: String, rightStrand: String): Int {
        if (leftStrand.length != rightStrand.length) {
            throw IllegalArgumentException("left and right strands must be of equal length")
        }
        var counter = 0
        leftStrand.zip(rightStrand) {l,r -> if(l != r) counter ++}
        return counter
    }
}
