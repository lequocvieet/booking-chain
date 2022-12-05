const { expect, use } = require('chai')
const { solidity } = require('ethereum-waffle')
const { ethers, upgrades } = require('hardhat')
const { convertBigNumberArray } = require('../helper-functions')
require('@nomiclabs/hardhat-waffle')
const { signWithdrawToken, signWithdrawNFT } = require('../helper-sign')

describe('Test', function () {
  beforeEach(async function () {
    ;[owner, signer, fakeSigner, user1, user2] = await ethers.getSigners()

    nonce = 1
    userId = []
    userId[user1] = '1'
    userId[user2] = '2'

    E721A = await hre.ethers.getContractFactory('E721A')
    e721A = await E721A.deploy('huyvip', 'HUY', '')
    await e721A.deployed()

    GToken = await hre.ethers.getContractFactory('GToken')
    gToken = await GToken.deploy('huyvip', 'HUY', '10000')
    await gToken.deployed()

    RToken = await hre.ethers.getContractFactory('RToken')
    rToken = await RToken.deploy('huyvip', 'HUY', '10000')
    await rToken.deployed()

    E1155 = await hre.ethers.getContractFactory('E1155')
    e1155 = await E1155.deploy('ipfs://')
    await e1155.deployed()
    //deploy
    GameService = await hre.ethers.getContractFactory('GameService')

    gameName = 'oke'
    gameService = await GameService.deploy(
      gameName,
      process.env.PUBLIC_KEY_BACKEND,
    )
    gameService.attachAddress(
      [rToken.address],
      [gToken.address],
      [e721A.address],
      [e1155.address],
    )

    await gameService.deployed()

    // Set approval
    maxUnit256 =
      '0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff'

    await e721A.connect(user1).setApprovalForAll(gameService.address, true)
    await e721A.connect(user2).setApprovalForAll(gameService.address, true)

    await e1155.connect(user1).setApprovalForAll(gameService.address, true)
    await e1155.connect(user2).setApprovalForAll(gameService.address, true)

    await gToken.connect(user1).approve(gameService.address, maxUnit256)
    await gToken.connect(user2).approve(gameService.address, maxUnit256)

    await rToken.connect(user1).approve(gameService.address, maxUnit256)
    await rToken.connect(user2).approve(gameService.address, maxUnit256)

    //set MINTER_ROLE for gameService.address
    await rToken.grantRole(
      ethers.utils.keccak256(ethers.utils.toUtf8Bytes('MINTER_ROLE')),
      gameService.address,
    )
    await rToken.grantRole(
      ethers.utils.keccak256(ethers.utils.toUtf8Bytes('BURNER_ROLE')),
      gameService.address,
    )
    await gToken.grantRole(
      ethers.utils.keccak256(ethers.utils.toUtf8Bytes('BURNER_ROLE')),
      gameService.address,
    )
    await e1155.grantRole(
      ethers.utils.keccak256(ethers.utils.toUtf8Bytes('MINTER_ROLE')),
      gameService.address,
    )
    // add minting permission for gameService.address

    await e721A.grantRole(
      ethers.utils.keccak256(ethers.utils.toUtf8Bytes('MINTER_ROLE')),
      gameService.address,
    )

    await e721A.mint(user1.address, 3)
    await e721A.mint(gameService.address, 3)

    await gToken.transfer(user1.address, ethers.utils.parseEther('1.0'))
    await rToken.transfer(user1.address, ethers.utils.parseEther('1.0'))

    await gToken.transfer(gameService.address, ethers.utils.parseEther('1.0'))
    await rToken.transfer(gameService.address, ethers.utils.parseEther('1.0'))

    await gameService.addValidator(signer.address)

    //fake e721A contract
    WrongNFTAddress = await hre.ethers.getContractFactory('E721A')
    wrongNFTAddress = await WrongNFTAddress.deploy('huyvip', 'HUY', '10000')
    await wrongNFTAddress.deployed()
  })

  it('1.1.1. Should withdrawToken using transfer TMR', async function () {
  
    tx = await gameService
      .connect(user1)
      .withdrawToken(
        [sign],
        nonce,
        gToken.address,
        ethers.utils.parseEther('0.1'),
      )

    await expect(tx)
      .to.emit(gameService, 'WithdrawToken')
      .withArgs(
        user1.address,
        1,
        gToken.address,
        ethers.utils.parseEther('0.1'),
      )

    nonce++

    sign = await signWithdrawToken(
      gameName,
      6996,
      gameService.address,
      nonce,
      user1.address,
      rToken.address,
      ethers.utils.parseEther('0.2'),
    )

    tx = await gameService.connect(user1).withdrawToken(
      [sign],
      nonce,

      rToken.address,
      ethers.utils.parseEther('0.2'),
    )

    await expect(tx).to.emit(gameService, 'WithdrawToken').withArgs(
      user1.address,
      2,

      rToken.address,
      ethers.utils.parseEther('0.2'),
    )

    expect(await gToken.balanceOf(user1.address)).to.equal(
      ethers.utils.parseEther('1.1'),
    )
    expect(await rToken.balanceOf(user1.address)).to.equal(
      ethers.utils.parseEther('1.2'),
    )

    expect(await gToken.balanceOf(gameService.address)).to.equal(
      ethers.utils.parseEther('0.9'),
    )
    expect(await rToken.balanceOf(gameService.address)).to.equal(
      ethers.utils.parseEther('0.8'),
    )
  })
  it('1.1.2. Should withdrawToken using minting TMR', async function () {
    sign = await signWithdrawToken(
      gameName,
      6996,
      gameService.address,
      nonce,
      user1.address,
      rToken.address,
      ethers.utils.parseEther('1.2'),
    )

    tx = await gameService
      .connect(user1)
      .withdrawToken(
        [sign],
        nonce,
        rToken.address,
        ethers.utils.parseEther('1.2'),
      )

    await expect(tx)
      .to.emit(gameService, 'WithdrawToken')
      .withArgs(
        user1.address,
        1,
        rToken.address,
        ethers.utils.parseEther('1.2'),
      )

    expect(await rToken.balanceOf(user1.address)).to.equal(
      ethers.utils.parseEther('2.2'),
    )

    expect(await rToken.balanceOf(gameService.address)).to.equal(
      ethers.utils.parseEther('1'), //minting so will not effect balance
    )
  })

  it('1.2. Should NOT withdrawToken because not enough balance', async function () {
    sign = await signWithdrawToken(
      gameName,
      6996,
      gameService.address,
      nonce,
      user2.address,

      gToken.address,
      ethers.utils.parseEther('2'), //too much
    )
    await expect(
      gameService.connect(user2).withdrawToken(
        [sign],
        nonce,

        gToken.address,
        ethers.utils.parseEther('2'), //too much
      ),
    ).to.be.revertedWith('ERC20: transfer amount exceeds balance')

    expect(await gToken.balanceOf(user2.address)).to.equal(
      ethers.utils.parseEther('0'),
    )
    expect(await rToken.balanceOf(user2.address)).to.equal(
      ethers.utils.parseEther('0'),
    )
  })

  it('1.3. Should not withdrawToken because invalid address/amount', async function () {
    sign = await signWithdrawToken(
      gameName,
      6996,
      gameService.address,
      nonce,
      user1.address,
      gToken.address,
      0,
    )

    await expect(
      gameService
        .connect(user1)
        .withdrawToken([sign], nonce, gToken.address, 0),
    ).to.be.revertedWith('withdrawToken: invalid balance')

    sign = await signWithdrawToken(
      gameName,
      6996,
      gameService.address,
      nonce,
      user1.address,
      wrongNFTAddress.address,
      ethers.utils.parseEther('0.2'),
    )

    await expect(
      gameService.connect(user1).withdrawToken(
        [sign],
        nonce,

        wrongNFTAddress.address,
        ethers.utils.parseEther('0.2'),
      ),
    ).to.be.revertedWith('GameService: Not our e20 address')
  })

  it('2.1. Should depositToken', async function () {
    tx = await gameService
      .connect(user1)
      .depositToken(gToken.address, ethers.utils.parseEther('0.1'))

    await expect(tx)
      .to.emit(gameService, 'DepositToken')
      .withArgs(user1.address, gToken.address, ethers.utils.parseEther('0.1'))

    tx = await gameService
      .connect(user1)
      .depositToken(rToken.address, ethers.utils.parseEther('0.2'))

    await expect(tx)
      .to.emit(gameService, 'DepositToken')
      .withArgs(user1.address, rToken.address, ethers.utils.parseEther('0.2'))

    expect(await gToken.balanceOf(user1.address)).to.equal(
      ethers.utils.parseEther('0.9'),
    )
    expect(await rToken.balanceOf(user1.address)).to.equal(
      ethers.utils.parseEther('0.8'),
    )

    expect(await gToken.balanceOf(gameService.address)).to.equal(
      ethers.utils.parseEther('1.1'),
    )
    expect(await rToken.balanceOf(gameService.address)).to.equal(
      ethers.utils.parseEther('1.2'),
    )
  })
  it('2.2. Should NOT depositToken because not enough balance', async function () {
    await expect(
      gameService
        .connect(user2)
        .depositToken(gToken.address, ethers.utils.parseEther('0.1')),
    ).to.be.revertedWith('ERC20: transfer amount exceeds balance')

    expect(await gToken.balanceOf(user2.address)).to.equal(
      ethers.utils.parseEther('0'),
    )
    expect(await rToken.balanceOf(user2.address)).to.equal(
      ethers.utils.parseEther('0'),
    )
  })
  it('2.3. Should NOT depositToken because invalid address and amount', async function () {
    await expect(
      gameService
        .connect(user1)
        .depositToken(wrongNFTAddress.address, ethers.utils.parseEther('0.1')),
    ).to.be.revertedWith('GameService: Not our e20 address')

    await expect(
      gameService.connect(user1).depositToken(gToken.address, 0),
    ).to.be.revertedWith('depositToken: invalid balance')
  })
  it('3.1. Should withdrawNFT', async function () {
    localIds = [
      ethers.utils.hexZeroPad(ethers.utils.formatBytes32String('10'), 32),
      ethers.utils.hexZeroPad(ethers.utils.formatBytes32String('11'), 32),
    ]

    //Character
    sign = await signWithdrawNFT(
      gameName,
      6996,
      gameService.address,
      nonce,
      user1.address,
      localIds,
      e721A.address,
    )
    tx = await gameService
      .connect(user1)
      .withdrawNFT([sign], nonce, localIds, e721A.address)

    await expect(tx)
      .to.emit(gameService, 'WithdrawNFT')
      .withArgs(user1.address, 1, localIds, e721A.address, 6)

    expect(await e721A.balanceOf(user1.address)).to.equal(5)
    expect(await e721A.balanceOf(gameService.address)).to.equal(3)

    expect(await e721A.tokensOfOwner(user1.address)).to.eql(
      convertBigNumberArray([0, 1, 2, 6, 7]),
    )
  })

  it('3.3. Should NOT withdrawNFT because wrong e721A address', async function () {
    localIds = [
      ethers.utils.hexZeroPad(ethers.utils.formatBytes32String('10'), 32),
      ethers.utils.hexZeroPad(ethers.utils.formatBytes32String('11'), 32),
    ]

    sign = await signWithdrawNFT(
      gameName,
      6996,
      gameService.address,
      nonce,
      user1.address,
      localIds,
      wrongNFTAddress.address,
    )
    await expect(
      gameService
        .connect(user1)
        .withdrawNFT([sign], nonce, localIds, wrongNFTAddress.address),
    ).to.be.revertedWith('GameService: Not our e721 address')
  })

  it('4.1. Should depositNFT character', async function () {
    //Character
    tx = await gameService.connect(user1).depositNFT(e721A.address, [0, 1])

    await expect(tx)
      .to.emit(gameService, 'DepositNFT')
      .withArgs(user1.address, e721A.address, [0, 1])

    expect(await e721A.balanceOf(user1.address)).to.equal(1)
    expect(await e721A.balanceOf(gameService.address)).to.equal(5)

    expect(await e721A.tokensOfOwner(user1.address)).to.eql(
      convertBigNumberArray([2]),
    )

    expect(await e721A.tokensOfOwner(gameService.address)).to.eql(
      convertBigNumberArray([0, 1, 3, 4, 5]),
    )
  })

  it('4.2. Should NOT depositNFT because gameService does not have that e721A ID', async function () {
    await expect(
      gameService.connect(user1).depositNFT(e721A.address, [4, 5]),
    ).to.be.revertedWith('TransferFromIncorrectOwner')

    await expect(
      gameService.connect(user1).depositNFT(e721A.address, [9, 10]),
    ).to.be.revertedWith('OwnerQueryForNonexistentToken')
  })

  it('4.3. Should NOT depositNFT because wrong e721A address', async function () {
    await expect(
      gameService.connect(user1).depositNFT(wrongNFTAddress.address, [4, 5]),
    ).to.be.revertedWith('GameService: Not our e721 address')
  })

  // it('6.1. Should burnToken', async function () {
  //   tx = await gameService
  //     .connect(owner)
  //     .burnToken(ethers.utils.parseEther('0.1'), ethers.utils.parseEther('0.2'))

  //   expect(tx)
  //     .to.emit(gameService, 'BurnToken')
  //     .withArgs(ethers.utils.parseEther('0.1'), ethers.utils.parseEther('0.2'))

  //   expect(await gToken.balanceOf(gameService.address)).to.equal(
  //     ethers.utils.parseEther('0.9'),
  //   )
  //   expect(await rToken.balanceOf(gameService.address)).to.equal(
  //     ethers.utils.parseEther('0.8'),
  //   )
  // })
  // it('6.2. Should not burnToken because not admin', async function () {
  //   await expect(
  //     gameService
  //       .connect(user1) //not admin
  //       .burnToken(
  //         ethers.utils.parseEther('0.1'),
  //         ethers.utils.parseEther('0.2'),
  //       ),
  //   ).to.be.revertedWith(
  //     'AccessControl: account ' +
  //       user1.address.toLowerCase() +
  //       ' is missing role 0x0000000000000000000000000000000000000000000000000000000000000000',
  //   )
  // })
  // it('7.1. Should burnNFT', async function () {
  //   tx = await gameService.connect(owner).burnNFT(e721A.address, [3])
  //   expect(tx).to.emit(gameService, 'BurnNFT').withArgs(e721A.address, [3])

  //   tx = await gameService.connect(owner).burnNFT(e721A.address, [4])
  //   expect(tx).to.emit(gameService, 'BurnNFT').withArgs(e721A.address, [4])

  //   expect(await e721A.tokensOfOwner(gameService.address)).to.eql(
  //     convertBigNumberArray([4, 5]),
  //   )
  //   expect(await e721A.tokensOfOwner(gameService.address)).to.eql(
  //     convertBigNumberArray([3, 5]),
  //   )
  // })
  // it('7.2. Should not burnNFT because not having e721A but have permission to burn', async function () {
  //   await expect(
  //     gameService.connect(owner).burnNFT(e721A.address, [0, 1]),
  //   ).to.be.revertedWith('gameService/burnNFT: not the owner')
  // })
  it('8.1. Should withdraw max is 5NFT/time', async function () {
    localIds = [
      ethers.utils.hexZeroPad(ethers.utils.formatBytes32String('10'), 32),
      ethers.utils.hexZeroPad(ethers.utils.formatBytes32String('11'), 32),
      ethers.utils.hexZeroPad(ethers.utils.formatBytes32String('12'), 32),
      ethers.utils.hexZeroPad(ethers.utils.formatBytes32String('13'), 32),
      ethers.utils.hexZeroPad(ethers.utils.formatBytes32String('14'), 32),
      ethers.utils.hexZeroPad(ethers.utils.formatBytes32String('15'), 32),
    ]

    sign = await signWithdrawNFT(
      gameName,
      6996,
      gameService.address,
      nonce,
      user1.address,
      localIds,
      e721A.address,
    )
    await expect(
      gameService.connect(user1).withdrawNFT(
        [sign],
        nonce,

        localIds,
        e721A.address,
      ),
    ).to.be.revertedWith('GameService: Invalid array')

    sign = await signWithdrawNFT(
      gameName,
      6996,
      gameService.address,
      nonce,
      user1.address,
      [],
      e721A.address,
    )
    await expect(
      gameService.connect(user1).withdrawNFT([sign], nonce, [], e721A.address),
    ).to.be.revertedWith('GameService: Invalid array')
  })
  it('8.2. Should deposit max is 5NFT/time', async function () {
    localIds = [
      ethers.utils.hexZeroPad(ethers.utils.formatBytes32String('6'), 32),
      ethers.utils.hexZeroPad(ethers.utils.formatBytes32String('7'), 32),
      ethers.utils.hexZeroPad(ethers.utils.formatBytes32String('8'), 32),
    ]

    sign = await signWithdrawNFT(
      gameName,
      6996,
      gameService.address,
      nonce,
      user1.address,
      localIds,
      e721A.address,
    )

    gameService
      .connect(user1)
      .withdrawNFT([sign], nonce, localIds, e721A.address)
    // user1 has 6 e721A character now

    await expect(
      gameService.connect(user1).depositNFT(e721A.address, [0, 1, 2, 6, 7, 8]),
    ).to.be.revertedWith('GameService: Invalid array')

    await expect(
      gameService.connect(user1).depositNFT(e721A.address, []),
    ).to.be.revertedWith('GameService: Invalid array')
  })

  // it.only('8.4. Testing max admin burn', async function () {
  //   await e721A.mint(gameService.address, 1000)

  //   await gameService
  //     .connect(owner)
  //     .burnNFT(e721A.address, Array.from(Array(1000).keys()))
  // })
  it('9.1. Should call ERC721AURI and ERC721AQueryable and ERC721ABurnable', async function () {
    expect(await e721A.tokenURI(0)).to.equal('')

    await e721A.connect(user1).burn(0)
    expect(await e721A.tokensOfOwner(user1.address)).to.eql(
      convertBigNumberArray([1, 2]),
    )
  })

  it('9.2. Should change baseURI', async function () {
    expect(await e721A.tokenURI(0)).to.equal('')

    await e721A.changeBaseURI('blablabla.com/1/')

    expect(await e721A.tokenURI(0)).to.equal('blablabla.com/1/0')
  })
})