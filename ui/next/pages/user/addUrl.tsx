// @ts-nocheck

import React, { useState } from 'react'
import TextField from '@mui/material/TextField'
import Button from '@mui/material/Button'
import Snackbar from '@mui/material/Snackbar'
import Alert from '@mui/material/Alert'
import Paper from '@mui/material/Paper'
import Typography from '@mui/material/Typography'
import IconButton from '@mui/material/IconButton'
import FileCopyIcon from '@mui/icons-material/FileCopy'
import Grid from '@mui/material/Grid'
import { CopyToClipboard } from 'react-copy-to-clipboard'
import Link from '@mui/material/Link'
import { Layout } from 'components'
import withAuthSync from 'components/Private'

function AddUrl() {
  const [open, setOpen] = useState(false)
  const classes = {}

  const [url, setURL] = useState({
    url: '',
  })

  const [response, setResponse] = useState({
    type: '',
    message: '',
    hash: '',
  })

  const handleChange = (e) =>
    setURL({ ...url, [e.target.name]: e.target.value })

  const handleClose = (event, reason) => {
    if (reason === 'clickaway') {
      return
    }

    setOpen(false)
  }

  const handleSubmit = async (e) => {
    e.preventDefault()
    try {
      // TODO: use store.actions
      const res = await fetch(`/api/link`, {
        method: 'POST',
        body: JSON.stringify(url),
        headers: { 'Content-Type': 'application/json' },
      })
      const json = await res.json()

      if (res.status === 201) {
        setResponse({
          type: 'success',
          message: 'Success add your link.',
          hash: json.hash,
        })
      } else {
        setResponse({
          type: 'error',
          message: json.error,
          hash: '',
        })
      }

      setOpen(true)
    } catch (error) {
      console.error('An error occurred', error) // eslint-disable-line
      setResponse({
        type: 'error',
        message: 'An error occured while submitting the form',
      })
      setOpen(true)
    }
  }

  return (
    <Layout>
      <Grid
        container
        direction="column"
        justifyContent="space-around"
        alignItems="center"
        className={classes.root}
      >
        <Paper className={classes.paper}>
          <form
            autoComplete="off"
            onSubmit={handleSubmit}
            className={classes.form}
          >
            <TextField label="Your URL" name="url" onChange={handleChange} />
            <TextField
              label="Describe"
              name="describe"
              onChange={handleChange}
            />
            <Button variant="contained" color="primary" type="submit">
              Add
            </Button>
          </form>
        </Paper>

        {response.type !== '' && response.type !== 'error' && (
          <Paper elevation={3} className={classes.paper}>
            <Typography variant="p" component="p">
              Your link: &nbsp;
              <Link
                href={`/s/${response.hash}`}
                target="_blank"
                rel="noopener"
                variant="body2"
                underline="hover"
              >
                {window.location.host}/s/{response.hash}
              </Link>
              <CopyToClipboard
                text={`${window.location.host}/s/${response.hash}`}
                onCopy={() => {
                  setResponse({
                    type: 'success',
                    message: 'Success copy your link.',
                    hash: response.hash,
                  })
                }}
              >
                <IconButton aria-label="copy" color="secondary" size="large">
                  <FileCopyIcon />
                </IconButton>
              </CopyToClipboard>
            </Typography>
          </Paper>
        )}

        <Snackbar
          anchorOrigin={{
            vertical: 'bottom',
            horizontal: 'left',
          }}
          open={open}
          autoHideDuration={6000}
          onClose={handleClose}
        >
          <Alert onClose={handleClose} severity={response.type}>
            {response.message}
          </Alert>
        </Snackbar>
      </Grid>
    </Layout>
  )
}

export default withAuthSync(AddUrl)
