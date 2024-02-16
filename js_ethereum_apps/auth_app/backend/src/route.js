const express = require('express')
const router = express.Router()
const authController = require('./controller')


function requireAuth(req, res, next) {
    if (!req.session.siwe) {
        res.status(401).json({ message: 'You have to first sign_in' })
        return
    }

    next()
}

router.get('/api/nonce', authController.Nonce)
router.post('/api/verify', authController.Verify)
router.get('/api/validate', requireAuth, authController.Validate)

module.exports = router
