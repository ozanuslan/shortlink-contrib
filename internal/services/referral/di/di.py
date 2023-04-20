"""Containers module."""

from dependency_injector import containers, providers

from usecases.crud_referral.crud import CRUDReferralService
from usecases.use_referral.use import UseReferralService

from infrastructure.repository.redis import Repository
from di.logger.logger import logger

class Core(containers.DeclarativeContainer):
  config = providers.Configuration("config")

  logger = providers.Resource(
    logger,
  )

class Application(containers.DeclarativeContainer):
  config = providers.Configuration()

  redis = providers.Singleton(
    Repository,
    host=config.redis_host,
  )

  core = providers.Container(
    Core,
    config=config.core,
  )

  referral_service = providers.Factory(
    CRUDReferralService,
    redis=redis,
  )

  use_service = providers.Factory(
    UseReferralService,
    redis=redis,
  )
