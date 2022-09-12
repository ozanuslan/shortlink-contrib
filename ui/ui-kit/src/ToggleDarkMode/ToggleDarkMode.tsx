import { useTheme as nextUseTheme } from 'next-themes'
import React, { useState, useContext, useEffect } from 'react'
import { ColorModeContext } from './ColorModeContext'

// @ts-ignore
import './styles.css'

export const ToggleDarkMode = () => {
  // @ts-ignore
  const { setTheme } = nextUseTheme()
  const [mounted, setMounted] = useState(false)

  useEffect(() => {
    setMounted(true)
  }, [])

  // @ts-ignore
  const { darkMode, setDarkMode } = useContext(ColorModeContext)

  // @ts-ignore
  const onChange = () => {
    setDarkMode(!darkMode)
    setTheme(darkMode ? 'light' : 'dark')
  }

  if (!mounted) return null

  return (
    <div className="toggleWrapper">
      <input
        type="checkbox"
        className="dn"
        id="dn"
        onChange={onChange}
        checked={!darkMode}
      />
      <label htmlFor="dn" className="toggle">
        <span className="toggle__handler">
          <span className="crater crater--1" />
          <span className="crater crater--2" />
          <span className="crater crater--3" />
        </span>
        <span className="star star--1" />
        <span className="star star--2" />
        <span className="star star--3" />
        <span className="star star--4" />
        <span className="star star--5" />
        <span className="star star--6" />
      </label>
    </div>
  )
}

export default ToggleDarkMode
