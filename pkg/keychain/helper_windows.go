// Copyright (c) 2024 Proton AG
//
// This file is part of Proton Mail Bridge.
//
// Proton Mail Bridge is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Proton Mail Bridge is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with Proton Mail Bridge. If not, see <https://www.gnu.org/licenses/>.

package keychain

import (
	"github.com/docker/docker-credential-helpers/credentials"
	"github.com/docker/docker-credential-helpers/wincred"
	"github.com/sirupsen/logrus"
)

const WindowsCredentials = "windows-credentials"

func listHelpers(_ bool) (Helpers, string) {
	helpers := make(Helpers)
	// Windows always provides a keychain.
	if isUsable(newWinCredHelper("")) {
		helpers[WindowsCredentials] = newWinCredHelper
		logrus.WithField("keychain", "WindowsCredentials").Info("Keychain is usable.")
	} else {
		logrus.WithField("keychain", "WindowsCredentials").Debug("Keychain is not available.")
	}
	// Use WindowsCredentials by default.
	return helpers, WindowsCredentials
}

func newWinCredHelper(string) (credentials.Helper, error) {
	return &wincred.Wincred{}, nil
}
